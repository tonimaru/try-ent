package main

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tonimaru/try-ent/ent"
	"github.com/tonimaru/try-ent/ent/fulltxt"
	entgeo "github.com/tonimaru/try-ent/ent/geo"
	"github.com/tonimaru/try-ent/pkg/geo"
	"github.com/tonimaru/try-ent/pkg/must"
)

func main() {
	ctx := context.Background()
	client := must.Do1(ent.Open("mysql", "root:pw@tcp(localhost:3306)/ent"))
	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}

	client.Geo.Delete().ExecX(ctx)

	client.Geo.CreateBulk(
		client.Geo.Create().SetPoint(&geo.Point{139.581973, 35.567485}),
		client.Geo.Create().SetPoint(&geo.Point{139.681973, 35.667485}),
		client.Geo.Create().SetPoint(&geo.Point{139.781973, 35.767485}),
	).ExecX(ctx)

	lon := 139.681974
	lat := 35.667486
	distanceExpr := sql.Expr("ST_Distance_Sphere(ST_SRID(point, 4326),ST_SRID(Point(?,?), 4326))", lon, lat)

	t := sql.Table(entgeo.Table)

	sub := sql.
		Select(t.C(entgeo.FieldID)).
		AppendSelectExprAs(distanceExpr, "distance").
		From(t).
		As("sub")

	geos := client.Geo.Query().
		Where(func(s *sql.Selector) {
			s.Join(sub).On(t.C(entgeo.FieldID), sub.C(entgeo.FieldID))
			s.AppendSelect(sub.C("distance"))
			s.Where(sql.GTE(sub.C("distance"), 1))
		}).
		AllX(ctx)
	for _, geo := range geos {
		fmt.Println(geo)
		fmt.Println(geo.Value("distance"))
	}

	client.Fulltxt.Delete().ExecX(ctx)

	// https://dev.mysql.com/doc/refman/8.0/en/fulltext-boolean.html
	client.Fulltxt.CreateBulk(
		client.Fulltxt.Create().SetTxt("This database tutorial ..."),
		client.Fulltxt.Create().SetTxt("After you went through a ..."),
		client.Fulltxt.Create().SetTxt("In this database tutorial ..."),
		client.Fulltxt.Create().SetTxt("When comparing databases ..."),
		client.Fulltxt.Create().SetTxt("When configured properly, MySQL ..."),
		client.Fulltxt.Create().SetTxt("database database database"),
		client.Fulltxt.Create().SetTxt("1. Never run mysqld as root. 2. ..."),
		client.Fulltxt.Create().SetTxt("MySQL fulltext indexes use a .."),
	).ExecX(ctx)

	txts := client.Fulltxt.Query().
		Where(func(s *sql.Selector) {
			s.
				Select(fulltxt.FieldID, fulltxt.FieldTxt).
				AppendSelectAs("MATCH (txt) AGAINST ('database' IN BOOLEAN MODE)", "score")
		}).AllX(ctx)
	for _, txt := range txts {
		fmt.Println(txt)
		fmt.Println(txt.Value("score"))
	}
}

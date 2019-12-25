package main

import (
	"encoding/json"
	"fmt"
	rp "thingz-server/rule/proto"
	tp "thingz-server/thing/proto"
)

// RuleM .
type RuleM struct {
	*rp.Rule
	DType []string `json:"dgraph.type,omitempty"`
	Rules []*RuleM `json:"rules,omitempty"`
}

func main() {
	r := &rp.Rule{
		Unit:        int32(tp.Unit_BOOL),
		TriggerType: rp.TriggerType_COMMAND,
		Operation:   rp.Operation_AND,
		Rules: []*rp.Rule{{
			Unit:        int32(tp.Unit_BOOL),
			TriggerType: rp.TriggerType_COMMAND,
			Operation:   rp.Operation_OR,
			Rules: []*rp.Rule{{
				Unit:        int32(tp.Unit_NUMBER),
				TriggerType: rp.TriggerType_COMMAND,
				Operation:   rp.Operation_EQUAL,
				Thing:       "100000",
				Channel:     "sensor",
				FloatValue:  18.0,
			}, {
				Unit:        int32(tp.Unit_NUMBER),
				TriggerType: rp.TriggerType_COMMAND,
				Operation:   rp.Operation_EQUAL,
				Thing:       "100000",
				Channel:     "sensor",
				FloatValue:  20.0,
			}},
		}, {
			Unit:        int32(tp.Unit_NUMBER),
			TriggerType: rp.TriggerType_COMMAND,
			Operation:   rp.Operation_EQUAL,
			Thing:       "100001",
			Channel:     "sensor",
			FloatValue:  18.0,
		}},
	}

	r.Evaluate(map[string]tp.Thing{
		"100000": tp.Thing{
			Id: "100000",
			Channels: []*tp.Channel{{
				Id:         "sensor",
				Unit:       tp.Unit_NUMBER,
				FloatValue: 18.0,
			}},
		},
		"100001": tp.Thing{
			Id: "100001",
			Channels: []*tp.Channel{{
				Id:         "sensor",
				Unit:       tp.Unit_NUMBER,
				FloatValue: 20.0,
			}},
		},
	})

	b, _ := json.MarshalIndent(r, "", "	")
	fmt.Println(string(b))

	rr := &rp.Rule{}

	json.Unmarshal(b, rr)

	vals := map[string]bool{}

	json.Unmarshal([]byte(rr.Vals), &vals)
	b, _ = json.MarshalIndent(vals, "", "	")

	fmt.Printf("%+v\n", string(b))

	fmt.Println(r.GetAllThings())

	fmt.Printf("%+v\n", r)

	rma := r.GetWithDType()

	b, _ = json.MarshalIndent(rma, "", "	")
	fmt.Println(string(b))
	json.Unmarshal(b, rma)
	fmt.Printf("%+v\n", rma.Thing)

	fmt.Println("------------")

	newR := &rp.Rule{
		Id: "1",
		TriggerCommands: []*rp.TriggerCommand{{
			Thing: "1",
			Channels: []*rp.Channel{{
				Id:         "sensor",
				FloatValue: 8.5,
				Unit:       1,
			}},
		}},
		Rules: []*rp.Rule{{
			Id: "2",
		}},
	}

	newR.AssignRootID(newR.Id)

	b, _ = json.MarshalIndent(newR, "", "	")
	fmt.Println(string(b))
	/* type School struct {
		Name  string   `json:"name,omitempty"`
		DType []string `json:"dgraph.type,omitempty"`
	}

	type loc struct {
		Type   string    `json:"type,omitempty"`
		Coords []float64 `json:"coordinates,omitempty"`
	}

	// If omitempty is not set, then edges with empty values (0 for int/float, "" for string, false
	// for bool) would be created for values not specified explicitly.

	type Person struct {
		Uid      string     `json:"uid,omitempty"`
		Name     string     `json:"name,omitempty"`
		Age      int        `json:"age,omitempty"`
		Dob      *time.Time `json:"dob,omitempty"`
		Married  bool       `json:"married,omitempty"`
		Raw      []byte     `json:"raw_bytes,omitempty"`
		Friends  []Person   `json:"friend,omitempty"`
		Location loc        `json:"loc,omitempty"`
		School   []School   `json:"school,omitempty"`
		DType    []string   `json:"dgraph.type,omitempty"`
	}

	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		age: int .
		married: bool .
		loc: geo .
		dob: datetime .

	type Person {
	  name: string
	  age: int
	  dob: string
	  married: bool
	  raw: string
	  friends: [Person]
	  loc: [Loc]
	  school: [Institution]
	 }

	type Loc {
	  type: string
	  coords: float
	 }

	type Institution {
	  name: string
	 }

	`

	ctx := context.Background()
	err = dg.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	dob := time.Date(1980, 01, 01, 23, 0, 0, 0, time.UTC)
	// While setting an object if a struct has a Uid then its properties in the graph are updated
	// else a new node is created.
	// In the example below new nodes for Alice, Bob and Charlie and school are created (since they
	// dont have a Uid).
	p := Person{
		Uid:     "_:alice",
		Name:    "Alice",
		Age:     26,
		Married: true,
		Location: loc{
			Type:   "Point",
			Coords: []float64{1.1, 2},
		},
		Dob: &dob,
		Raw: []byte("raw_bytes"),
		Friends: []Person{{
			Name: "Bob",
			Age:  24,
		}, {
			Name: "Charlie",
			Age:  29,
		}},
		School: []School{{
			Name: "Crown Public School",
		}},
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id": assigned.Uids["alice"]}
	q := `query Me($id: string){
		me(func: uid($id)) {
			name
			dob
			age
			loc
			raw_bytes
			married
			dgraph.type
			friend @filter(eq(name, "Bob")){
				name
				age
				dgraph.type
			}
			school {
				name
				dgraph.type
			}
		}
	}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	type Root struct {
		Me []Person `json:"me"`
	}

	var root Root
	err = json.Unmarshal(resp.Json, &root)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Me: %+v\n", r.Me)
	// R.Me would be same as the person that we set above.

	fmt.Println(string(resp.Json)) */
	// Output: {"me":[{"name":"Alice","dob":"1980-01-01T23:00:00Z","age":26,"loc":{"type":"Point","coordinates":[1.1,2]},"raw_bytes":"cmF3X2J5dGVz","married":true,"dgraph.type":["Person"],"friend":[{"name":"Bob","age":24,"dgraph.type":["Person"]}],"school":[{"name":"Crown Public School","dgraph.type":["Institution"]}]}]}
}

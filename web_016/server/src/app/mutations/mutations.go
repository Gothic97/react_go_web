package mutations
import (
	"github.com/graphql-go/graphql"
	fields "gothic/web_016/server/src/app/mutations/fields"
)
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": fields.CreateNotTodo,
	},
})
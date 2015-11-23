package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/ai","description":"Operations about AI\n"},{"path":"/webchat","description":"Operations about AI\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/ai":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/ai","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/:chat","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"ai chat","parameters":[{"paramType":"path","name":"chat","description":"\"Please input chat\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.AI","responseModel":""},{"code":403,"message":":chat is empty","responseModel":""}]}]}]},"/webchat":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/webchat","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"createUser","type":"","summary":"get webchat content","parameters":[{"paramType":"body","name":"body","description":"\"body for webchat content\"","dataType":"WebChat","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.WebChat.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"WebChat Check","responseMessages":[{"code":200,"message":":echostr","responseModel":""},{"code":403,"message":":echostr is empty","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	if beego.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocApi["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.Apis {
				a.Path = urlReplace(k + a.Path)
				v.Apis[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocApi[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}

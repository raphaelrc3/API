package rotas

import "net/http"

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		funcao: Controllers.CriarUsuario,
		RequerAutenticação: false,
    },
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		funcao: controllers.BuscarUsuarios,
		RequerAutenticação: false,
	},
	{
		URI:    "/usuarios/{usuarioID}",
		Metodo: http.MethodGet,
		funcao: controles.BuscarUsuario,
		RequerAutenticação: false,
	},
	{
		URI:    "/usuarios/{usuarioID}",
		Metodo: http.MethodPut,
		funcao: controllers.AtualizarUsuario
		RequerAutenticação: false,
	},
	{
		URI:    "/usuarios/{usuarioID}",
		Metodo: http.MethodDelete,
		funcao: Controllers.DeletarUsuario
		RequerAutenticação: false,
	},
	
}
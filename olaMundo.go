/*
 * "Olá Mundo!" em Go com os bindings do gotk3 para o toolkit
 * de interfaces gráficas GTK 3
 *
 * Documentação em: https://pkg.go.dev/github.com/gotk3/gotk3/gtk
 *
*/

package main

import (

	"os"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"

)

// ID da aplicação

const appId = "com.github.brendocosta.gotk3-olamundo"

func checarErros(e error) {

	if e != nil {

		// Exibe o erro

		log.Fatalln("Ocorreu um erro:", e)

	}

}

func main() {

	// Cria a instância da aplicação

	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	checarErros(err)

	/* Passo opcional:
	 * Função anônima para exibir uma mensagem
	 * quando o evento de inicialização da
	 * aplicação for disparado.
	*/

	application.Connect("startup", func() {

		log.Println("Inicializando a aplicação...")

	})

	/* Função anônima para o evento de ativação
	 * da aplicação. É a parte principal da
	 * execução do programa.
	*/

	application.Connect("activate", func() {

		log.Println("Aplicação inicializada!")

		/* Carrega o modelo da interface do Glade no
		 * builder a partir de um arquivo. Você pode
		 * substituir AddFromFile por AddFromString(string)
		 * caso queira carregar o modelo diretamente
		 * de uma string.
	    */

		builder, err := gtk.BuilderNew()
		checarErros(err)

		builder.AddFromFile("/home/neofox/Documentos/olaMundo.glade")

		// Pega o objeto com o ID "JanelaPrincipal"

		obj, err := builder.GetObject("JanelaPrincipal")
		checarErros(err)

		// Type assertion para o tipo *gtk.Window

		JanelaPrincipal, ok := obj.(*gtk.Window)
		if !ok { log.Fatalln("Ocorreu um erro: o objeto não encontrado ou não é uma janela!") }

		// Exibe a janela principal e todos os seus componentes

		JanelaPrincipal.ShowAll()

		// Adiciona a janela a instância da aplicação

		application.AddWindow(JanelaPrincipal)

		/* Mapeia os eventos do modelo para as funções do código
		 * Você também pode criar funções fora de main() e especificar
		 * apenas o nome dela sem parêntesis ao invés de utilizar
		 * uma função anônima.
	    */

		eventos := map[string] interface {} {

			"BotaoClicado": func () {

			    log.Println("O botão foi clicado!")

			    obj, err := builder.GetObject("JanelaDialogo")
	            checarErros(err)

	            JanelaDialogo, ok := obj.(*gtk.Window)
	            if !ok { log.Fatalln("Ocorreu um erro: o objeto não encontrado ou não é uma janela!") }

	            JanelaDialogo.ShowAll()

	            JanelaDialogo.HideOnDelete()

			},

		}

		builder.ConnectSignals(eventos)

	})

    application.Connect("shutdown", func () {

        log.Println("Aplicação encerrada!")

    })

	os.Exit(application.Run(os.Args))

}

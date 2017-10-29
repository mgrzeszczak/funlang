import React, {Component} from 'react'
import Console from 'react-console-component'
import 'react-console-component/main.css'

export default class WebConsole extends Component {
  constructor(props){
    super(props)
    this.props = props
    this.echo = (text) => {
      if (text == './funlang') {
        this.props.onExecute(this.console);
      } else if (text == 'help') {
        this.console.log("Available commands are: './funlang', 'help'");
      } else {
        this.console.log("Invalid command. Type 'help' to show available commands.")
      }
      this.console.return();
  	};
    this.promptLabel = () => {
  		return ">> ";
  	}

  }

  render(){
    return (
      <Console ref={ref => this.console = ref}
  			handler={this.echo}
  			promptLabel={this.promptLabel}
  			welcomeMessage={"Welcome to funlang playground console!\nType help to show available commands."}
  			autofocus={true}
  		/>
    )
  }
}

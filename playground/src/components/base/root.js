import React, {Component} from 'react'
import Editor from './editor'
import WebConsole from './web-console'

export default class Root extends Component {
  constructor(props){
    super(props)
  }

  onExecute(webconsole) {
    try {
      funlang.runScript(this.state.editor.code, webconsole.log);
    } catch (e) {
        webconsole.logX("error", "An error occurred:")
        webconsole.logX("error", e.message);
    }
  }

  onEditorChange(editorState){
    this.setState({
      editor:editorState
    });
  }

  render(){
    return (
        <div className="container">
          <div className="editor">
            <Editor onChange={this.onEditorChange.bind(this)}/>
          </div>
          <div className="console">
            <WebConsole onExecute={this.onExecute.bind(this)}/>
          </div>
        </div>
    )
  }
}

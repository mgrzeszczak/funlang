import React, {Component} from 'react'
import CodeMirror from 'react-codemirror'

export default class Editor extends Component {
  constructor(props){
    super(props);
    this.props = props;
    this.state = {
      code : "// type code here"
    };
    this.options = {
			lineNumbers: true,
      theme: 'monokai'
		};
  }

  updateCode(newCode) {
    this.setState({
      code: newCode,
    });
    this.props.onChange(this.state);
  }

  componentDidMount() {
    this.refs.editor.codeMirror.setSize("100%", "100%")
  }

  render(){
    return (
      <CodeMirror
        ref="editor"
        value={this.state.code}
        onChange={this.updateCode.bind(this)}
        options={this.options}
      />
    )
  }
}

import React,{Component} from 'react'
import {HashRouter,Route,Switch} from 'react-router-dom'
import Root from './components/base/root'
export default () => {
  return (
    <HashRouter>
      <div>
        <Switch>
          <Route path="/" component={Root}/>
        </Switch>
      </div>
    </HashRouter>
  )
}

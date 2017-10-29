import React,{Component} from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux'
import {createStore,applyMiddleware} from 'redux';
import ReduxPromise from 'redux-promise'

import App from './app'

const createStoreWithMiddleware = applyMiddleware(ReduxPromise)(createStore);

ReactDOM.render(
  <Provider store={createStoreWithMiddleware}>
    <div className="container">
      <App />
    </div>
  </Provider>,
  document.querySelector('.app')
);

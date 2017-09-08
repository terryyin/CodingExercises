import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';

let props = {rows: 150, cols: 150, seeds: 4000};
ReactDOM.render(<App {...props}/>, document.getElementById('root'));
registerServiceWorker();

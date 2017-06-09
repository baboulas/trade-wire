import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import EntryForm from './components/EntryForm';

import registerServiceWorker from './registerServiceWorker';
import './style/style.css';

// Import react router deps
import {Router, Route, IndexRoute, browserHistory } from 'react-router';
import {Provider} from 'react-redux';
import store, {history} from './store';

const router = {
    <provider store={store}>
    <Router history={browserHistory}>
        <Route path="/" component={App}>
        <IndexRoute component={EntryForm}></IndexRoute>
        </Route>     
    </Router>
    </provider>
}

ReactDOM.render(router, document.getElementById('root'));
registerServiceWorker();

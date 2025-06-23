import {useState} from 'react'; //unused currently, plan to build with state management later
import './assets/styles/style.css';

import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Layout from './components/Layout';
import Home from './pages/Home';
import AlbumView from './pages/AlbumView';
import Settings from './pages/Settings';

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
	  <Route path="/album/:albumName" element={<AlbumView />} />
          <Route path="/settings" element={<Settings />} />
        </Routes>
      </Layout>
    </Router>
  );
}

export default App;


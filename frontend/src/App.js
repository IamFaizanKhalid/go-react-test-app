import React from 'react';
import './App.css';
import HelloWorld from './components/HelloWorld';

function App() {
  return (
    <div id="app" className="App">
      <header className="App-header">
        <p>
          Welcome to your new <code>wails/react</code> project.
        </p>

        <HelloWorld />
      </header>
    </div>
  );
}

export default App;

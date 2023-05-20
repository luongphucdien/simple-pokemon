import React, {useState} from 'react';
import { test } from './API';
import { Login } from './Components/Login';
import { World } from './Components/World';
import { Test } from './Components/Test';

function App() {
  const [username, setUsername] = useState("")

  return (
    <>
      {/* <World/> */}
      {username ? <World/> : <Login onLoginSuccess={setUsername}/>}
      {/* <Test/> */}
    </>
  );
}

export default App;

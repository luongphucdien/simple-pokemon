import React, {useEffect, useState} from 'react';
import { Login } from './Components/Login';
import { World } from './Components/World';

function App() {
  const [username, setUsername] = useState("")

  return (
    <>
      {/* <World/> */}
      {username ? <World username={username} setUsername={setUsername}/> : <Login onLoginSuccess={setUsername}/>}
      {/* <Test/> */}
    </>
  );
}

export default App;

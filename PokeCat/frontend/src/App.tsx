import React, {useState} from 'react';
import { test } from './API';
import { Login } from './Components/Login';
import { World } from './Components/World';

function App() {
  const [isHidden, setIsHidden] = useState(false)

  return (
    <>
      {/* <World/> */}
      {isHidden ? <World/> : <Login onLoginSuccess={setIsHidden}/>}
    </>
  );
}

export default App;

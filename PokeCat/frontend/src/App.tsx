import React, {useState} from 'react';
import { test } from './API';
import { Login } from './Components/Login';

function App() {

  const [testData, setTestData] = useState()

  const OnClickHandler =  async () => {
    await test(setTestData)
  } 

  return (
    <>
      <Login/>
    </>
  );
}

export default App;

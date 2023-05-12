import { Button, TextField, Typography } from '@mui/material';
import React, {useState} from 'react';
import { test } from './API';
import axios from 'axios';

function App() {

  const [testData, setTestData] = useState()

  const OnClickHandler =  async () => {
    await test(setTestData)
    console.log(testData)
    // axios.get("http://localhost:8080/api/test").then(r => {
    //   console.log(r.data)
    // })
  } 

  return (
    <>
      <div>
        <Button 
          variant='contained'
          onClick={OnClickHandler}
        >
          Test
        </Button>

        <p>{testData}</p>
      </div>
    </>
  );
}

export default App;

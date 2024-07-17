import { useState } from 'react';
import { Button } from './components/ui/button';

function App() {
  const [name, setName] = useState('');

  return (
    <div id="App">
      <Button>
        Hola
      </Button>
    </div>
  );
}

export default App;

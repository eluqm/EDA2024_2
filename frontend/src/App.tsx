import { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { SearchSongInIndexInvert, SearchSongInTrie } from '../wailsjs/go/main/App';
import Browser from './components/browsers/Browser';
import { ResultType } from './types/result';

function App() {
  const [input, setInput] = useState('');

  const initResult: ResultType = {
    Songs: [],
    TimeLapse: 0,
    Size: 0,
  };

  const [resultInvertIndex, setResultInvertIndex] = useState<ResultType>(initResult);
  const [resultTrie, setResultTrie] = useState<ResultType>(initResult);

  const handleSearch = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (input.trim() === '') {
      setResultInvertIndex(initResult);
      setResultTrie(initResult);
      return;
    }

    const rTrie = await SearchSongInTrie(input) as ResultType;
    const rInvertIndex = await SearchSongInIndexInvert(input) as ResultType;

    setResultInvertIndex(rInvertIndex);
    setResultTrie(rTrie);
  };

  return (
    <div className="w-full flex flex-col gap-10 items-center justify-start py-10 px-6 flex-1">
      <form onSubmit={handleSearch} className="flex flex-col  w-full max-w-lg gap-2">
        <h1 className="text-2xl font-bold uppercase text-center w-full">
          Search for name of the song
        </h1>
        <div className="flex w-full gap-4">
          <Input onChange={(e) => setInput(e.target.value)} value={input} />
          <Button type="submit">Search</Button>
        </div>
      </form>
      <div className="w-full grid grid-cols-2 max-w-7xl border border-border rounded-xl min-h-[70vh]">
        <Browser
          result={resultInvertIndex.Songs}
          title="Invert index"
          timeLapse={resultInvertIndex.TimeLapse}
          size={resultInvertIndex.Size}
        />
        <Browser
          result={resultTrie.Songs}
          title="Trie"
          timeLapse={resultTrie.TimeLapse}
          size={resultTrie.Size}
        />
      </div>
    </div>
  );
}

export default App;

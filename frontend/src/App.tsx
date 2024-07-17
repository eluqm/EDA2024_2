import { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { SongType } from './types/song';
import { SearchSong } from '../wailsjs/go/main/App';
import Song from './components/ui/song';

function App() {
  const [search, setSearch] = useState('');
  const [result, setResult] = useState<SongType[]>([]);

  const handleSearch = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // si el string solo son espacios que no aga nada
    if (search.trim() === '') {
      setResult([]);
      return;
    }

    const r = await SearchSong(search) as SongType[];
    setResult(r);
  };

  return (
    <div className="w-full max-w-6xl flex flex-col gap-10 items-center justify-start py-10 px-6">
      <form onSubmit={handleSearch} className="flex flex-col w-full max-w-lg gap-2">
        <h1 className="text-2xl font-bold uppercase text-center w-full">
          Search for name of the song
        </h1>
        <div className="flex w-full gap-4">
          <Input onChange={(e) => setSearch(e.target.value)} value={search} />
          <Button type="submit">Search</Button>
        </div>
      </form>
      <div className="w-full">
        <div className="grid grid-cols-5 py-3 px-6 text-muted-foreground">
          <span className="col-span-2">
            Track Name
          </span>
          <span>
            Artist Name
          </span>
          <span>
            Track Id
          </span>
          <span className="flex justify-end items-center">
            Genre
          </span>
        </div>
        {result.map((song) => (
          <Song key={song.TrackId} song={song} />
        ))}
      </div>

    </div>
  );
}

export default App;

import { SongType } from '@/types/song';
import { convertirMilisegundos } from '@/lib/utils';
import { MinusIcon } from 'lucide-react';
import { Button } from './button';
import { usePlayList } from '../providers/PlayListProvider';

export default function PlayListSong({ song, index }: { song: SongType, index: number }) {
  const { removeSong } = usePlayList();

  return (
    <div className="flex border-t border-border py-3 px-6 gap-4 items-center">
      <span className="overflow-x-hidden text-nowrap w-full max-w-8">
        {index + 1}
      </span>
      <h1 className="overflow-x-hidden text-nowrap w-full">{song.TrackName}</h1>
      <h2 className="w-full max-w-64 overflow-x-hidden text-nowrap">{song.ArtistName}</h2>
      <h2 className="w-full max-w-24 text-center">{song.Popularity}</h2>
      <h2 className="w-full max-w-24 text-end">{song.Year}</h2>
      <h2 className="w-full max-w-24 flex justify-end items-center">{convertirMilisegundos(song.DurationMs)}</h2>
      <div className="w-full max-w-16 flex justify-end items-center">
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            removeSong(index);
          }}
        >
          <MinusIcon className="w-4" />
        </Button>
      </div>
    </div>
  );
}

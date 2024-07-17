import { SongType } from '@/types/song';

export default function Song({ song }: { song: SongType }) {
  return (
    <div className="grid grid-cols-6 border-t border-border py-3 px-6">
      <h1 className="col-span-4 overflow-hidden">{song.TrackName}</h1>
      <h2>{song.ArtistName}</h2>
      <h2 className="w-full flex justify-end items-center">{song.TrackId}</h2>
    </div>
  );
}

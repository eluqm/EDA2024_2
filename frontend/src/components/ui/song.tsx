import { SongType } from '@/types/song';

export default function Song({ song }: { song: SongType }) {
  return (
    <div className="grid grid-cols-5 border-t border-border py-3 px-6">
      <h1 className="col-span-2">{song.TrackName}</h1>
      <h2>{song.ArtistName}</h2>
      <h2>{song.TrackId}</h2>
      <h2 className="w-full flex justify-end items-center">{song.Genre}</h2>
    </div>
  );
}

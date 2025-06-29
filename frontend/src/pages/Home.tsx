import { useEffect, useState } from 'react';
import { EventsOn } from '../../wailsjs/runtime';
import AlbumCard from '../components/AlbumCard';
import '../assets/styles/card.css';
import { GetAllTracks } from '../../wailsjs/go/libmanager/Service';

type AlbumInfo = { albumName: string; artistName?: string; coverPath: string };
type Track = Awaited<ReturnType<typeof GetAllTracks>>[number];

const Home = () => {
  const [albums, setAlbums] = useState<AlbumInfo[]>([]);

  const loadLibrary = async () => {
    try {
      const tracks: Track[] = await GetAllTracks();
      const albumMap = new Map<string, AlbumInfo>();

      for (const track of tracks) {
        const album = track.album || 'Unknown Album';

        if (!albumMap.has(album)) {
          const coverPath = track.coverPath;

          albumMap.set(album, {
            albumName: album,
            artistName: track.artist || 'Unknown Artist',
            coverPath,
          });
        }
      }

      setAlbums(Array.from(albumMap.values()));
    } catch (err) {
      console.error('Failed to load tracks:', err);
    }
  };

  // wails event should trigger refresh
  useEffect(() => {
    loadLibrary();

    // event emitted from sidebar
    const unsubscribe = EventsOn("library_updated", loadLibrary);

    return () => unsubscribe();
  }, []);

  return (
    <div className="album-grid-container">
      {albums.length === 0 ? (
        <p style={{ color: 'white' }}>No albums found.</p>
      ) : (
        albums.map(album => (
          <AlbumCard
            key={album.albumName}
            albumName={album.albumName}
            artistName={album.artistName}
            coverPath={album.coverPath}
          />
        ))
      )}
    </div>
  );
};

export default Home;


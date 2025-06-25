import '../assets/styles/layout.css';
import { AddMusicFolder, OpenDirectorySelector, RescanLibrary } from '../../wailsjs/go/libmanager/Service';
import { EventsEmit } from '../../wailsjs/runtime';

export default function Sidebar() {
const handleAddMusic = async () => {
  try {
    const folder = await OpenDirectorySelector();
    if (!folder) return;

    const tracks = await AddMusicFolder(folder);
    alert(`Added ${tracks.length} tracks to library.`);
    await RescanLibrary();

    // home should reload albums after library update now
    EventsEmit("library_updated");
  } catch (error) {
    console.error("Error adding music folder:", error);
    alert("Failed to add music from folder.");
  }
};

  return (
    <aside className="sidebar">
      {/*Logo is a button to rescan aswell in case of non-scanning lib*/}
      <button className="logo" onClick={async() => {
	      await RescanLibrary();
      	      EventsEmit("library_updated");}}>
	      LOGO
      </button>
      <button className="add-button" onClick={handleAddMusic}>
        + Add Music
      </button>
    </aside>
  );
}


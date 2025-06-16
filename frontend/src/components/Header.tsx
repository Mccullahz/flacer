import "../assets/styles/layout.css";

type HeaderProps = {
  trackIcon?: string; // Default icon is music note, change later
  trackTitle?: string;
  trackArtist?: string;
  onPlayPause?: () => void;
  onSkipForward?: () => void;
  onSkipBackward?: () => void;
};

export default function Header({
  trackIcon = "🎵",
  trackTitle = " ",
  trackArtist = " ",
  onPlayPause,
  onSkipForward,
  onSkipBackward,
}: HeaderProps) {
  return (
    <header className="header">
      <div className="music-controls">
        <button className="control-button" onClick={onSkipBackward} aria-label="Rewind">⏮</button>
	<button className="control-button" onClick={onPlayPause} aria-label="Play/Pause">▶️</button>
        <button className="control-button" onClick={onSkipForward} aria-label="Forward">⏭</button>
      </div>

      <div className="now-playing">
        <span className="track-icon">{trackIcon}</span>
        <span className="track-title">{trackTitle}</span>
        <span className="track-artist">{trackArtist}</span>
      </div>

      <div className="header-actions">
      </div>
    </header>
  );
}

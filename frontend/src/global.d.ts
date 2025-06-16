export {};

declare global {
  interface Window {
    showOpenFilePicker?: (options?: any) => Promise<any>;

    backend: {
      Libmanager: {
        Service: {
          AddMusicFile(filePath: string): Promise<any>;
        };
      };
    };
  }
}


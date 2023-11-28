import React, { createContext, useContext, ReactNode, useState } from 'react';

interface SyncFileDataContextProps {
  synchronizedFiles: { filename: string, filepath: string }[] | null;
  updateSynchronizedFiles: (newFiles: { filename: string, filepath: string }[]) => void;
}

const DataContext = createContext<SyncFileDataContextProps | undefined>(undefined);

export const useSyncFileDataContext = () => {
  const context = useContext(DataContext);
  if (!context) {
    throw new Error('useDataContext must be used within a SyncFileDataProvider');
  }
  return context;
};

interface SyncFileDataProviderProps {
  children: ReactNode;
}

export const SyncFileDataProvider: React.FC<SyncFileDataProviderProps> = ({ children }) => {
  const [synchronizedFiles, setSynchronizedFiles] = useState<{ filename: string, filepath: string }[] | null>(null);

  const updateSynchronizedFiles = (newFiles: { filename: string, filepath: string }[]) => {
    setSynchronizedFiles(newFiles);
  };

  const contextValue: SyncFileDataContextProps = {
    synchronizedFiles,
    updateSynchronizedFiles,
  };

  return <DataContext.Provider value={contextValue}>{children}</DataContext.Provider>;
};

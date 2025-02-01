import React, { createContext, useState, useContext, ReactNode } from 'react';

export interface UserContextProps {   // Definimos las propiedades que tendrá el Usuario
  username: string;
  id: number;
  email: string;
  dayselected: string;
  isLoggedIn: boolean;
}

export const UserContext = createContext<UserContextProps | null>(null);


export const useUser = () => {
  const user = useContext(UserContext);
  if (!user) {
    throw new Error('useUser must be used within a UserProvider');
  }
  return user;
};


export default useUser;
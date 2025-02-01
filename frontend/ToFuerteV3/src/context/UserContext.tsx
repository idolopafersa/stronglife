import React, { createContext, useState, useContext, ReactNode } from 'react'

export interface UserContextProps {
  username: string
  id: number
  email: string
  dayselected: string
  isLoggedIn: boolean
}

interface UserContextType {
  currentUser: UserContextProps
  setCurrentUser: React.Dispatch<React.SetStateAction<UserContextProps>>
}

// Create Context with default value as null
export const UserContext = createContext<UserContextType | null>(null)

// Hook to use User Context
export const useUser = () => {
  const context = useContext(UserContext)
  if (!context) {
    throw new Error('useUser must be used within a UserProvider')
  }
  return context
}

// UserProvider Component
export const UserProvider = ({ children }: { children: ReactNode }) => {
  const [currentUser, setCurrentUser] = useState<UserContextProps>({
    username: '',
    id: 0,
    email: '',
    dayselected: '',
    isLoggedIn: false,
  })

  return (
    <UserContext.Provider value={{ currentUser, setCurrentUser }}>
      {children}
    </UserContext.Provider>
  )
}

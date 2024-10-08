import React from 'react'
import DateDisplay from '../components/DateDisplay'

const HomePage: React.FC = () => {
    return (
        <>
            <div className="flex items-center justify-center min-h-screen bg-gray-100">
                <h1 className="text-4xl font-bold text-blue-500">
                    Hello, Tailwind CSS with React!
                </h1>
            </div>
            <div style={{
                position: 'relative',
                width: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                flexDirection: 'column'
            }}>
                <h1 style={{ fontSize: '4em' }}>Hello world!</h1>
                <DateDisplay />
            </div>
        </>


    )
}

export default HomePage

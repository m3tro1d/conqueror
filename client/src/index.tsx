import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Dashboard from './components/Dashboard/Dashboard'
import LoginForm from './components/AuthForms/LoginForm/LoginForm'
import SignUpForm from './components/AuthForms/SignUpForm/SignUpForm'

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement,
)

const router = createBrowserRouter([
    {
        path: '/',
        element: <Dashboard />,
    },
    {
        path: '/login',
        element: <LoginForm />,
    },
    {
        path: '/signup',
        element: <SignUpForm />,
    },
])

root.render(
    <React.StrictMode>
        <RouterProvider router={router} />
    </React.StrictMode>,
)

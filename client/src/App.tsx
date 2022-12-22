import { BrowserRouter, Route, Routes } from 'react-router-dom'
import LoginPage from './components/AuthPages/LoginPage/LoginPage'
import SignUpPage from './components/AuthPages/SignUpPage/SignUpPage'
import Dashboard from './components/Dashboard/Dashboard'
import TasksPage from './components/TasksPage/TasksPage'
import useToken from './hooks/useToken'

function App() {
    const { token, setToken } = useToken()
    if (!token) {
        if (document.location.pathname == '/signup') {
            return <SignUpPage />
        }

        return <LoginPage setToken={setToken} />
    }

    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Dashboard />} />
                <Route path="/tasks" element={<TasksPage />} />
            </Routes>
        </BrowserRouter>
    )
}


export default App

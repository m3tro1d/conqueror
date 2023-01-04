import { BrowserRouter, Route, Routes } from 'react-router-dom'
import LoginPage from './components/AuthPages/LoginPage/LoginPage'
import SignUpPage from './components/AuthPages/SignUpPage/SignUpPage'
import Dashboard from './components/AppPages/Dashboard/Dashboard'
import TasksPage from './components/AppPages/TasksPage/TasksPage'
import useToken from './hooks/useToken'
import Header from './components/AppPages/common/Header/Header'
import Navigation from './components/AppPages/common/Navigation/Navigation'
import SubjectsPage from './components/AppPages/SubjectsPage/SubjectsPage'
import NotesPage from './components/AppPages/NotesPage/NotesPage'
import EditTaskPage from './components/AppPages/EditTaskPage/EditTaskPage'
import EditNotePage from './components/AppPages/EditNotePage/EditNotePage'

function App() {
    const { token, setToken } = useToken()
    if (!token) {
        if (document.location.pathname == '/signup') {
            return <SignUpPage />
        }

        return <LoginPage setToken={setToken} />
    }

    return (
        <div>
            <Header />
            <Navigation />
            <div>
                <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<Dashboard />} />

                        <Route path="/subjects" element={<SubjectsPage />} />

                        <Route path="/tasks" element={<TasksPage />} />
                        <Route path="/task/:id" element={<EditTaskPage />} />

                        <Route path="/notes" element={<NotesPage />} />
                        <Route path="/note/:id" element={<EditNotePage />} />
                    </Routes>
                </BrowserRouter>
            </div>
        </div>
    )
}


export default App

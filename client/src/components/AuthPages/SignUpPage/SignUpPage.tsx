import React, { FormEvent, useState } from 'react'
import styles from './SignUpPage.module.css'
import {userApi} from '../../../api/api'

function SignUpPage(): JSX.Element {
    const [login, setLogin] = useState('')
    const [password, setPassword] = useState('')

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()

        try {
            await userApi.signup({
                login,
                password,
            })
            document.location.href = '/'
        } catch (error) {
            alert('Failed to register.')
        }
    }

    return (
        <form
            className={styles.form}
            onSubmit={handleSubmit}
        >
            <h1 className={styles.header}>Sign up for Conqueror</h1>

            <input type="text"
                   name="login"
                   placeholder="Login"
                   className={styles.input}
                   onChange={e => setLogin(e.target.value)}
            />
            <br />

            <input type="password"
                   name="password"
                   placeholder="Password"
                   className={styles.input}
                   onChange={e => setPassword(e.target.value)}
            />
            <br />

            <div className={styles.buttonsContainer}>
                <button
                    className={styles.signupButton}
                >
                    Sign up
                </button>
            </div>
        </form>
    )
}

export default SignUpPage

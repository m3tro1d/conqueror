import React from 'react'
import styles from './LoginForm.module.css'

function LoginForm(): JSX.Element
{
    return (
        <form
            className={styles.form}
            onSubmit={event => event.preventDefault()}
        >
            <h1 className={styles.header}>Conqueror</h1>

            <input type="text" name="login" placeholder="Login" className={styles.input} />
            <br />

            <input type="password" name="password" placeholder="Password" className={styles.input} />
            <br />

            <div className={styles.buttonsContainer}>
                <button type="submit" className={styles.loginButton}>Login</button>
                <button
                    className={styles.signupButton}
                    onClick={event => {
                        event.preventDefault()
                        document.location.href = '/signup'
                    }}
                >
                    Sign up
                </button>
            </div>
        </form>
    )
}

export default LoginForm

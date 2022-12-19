import React from 'react';
import DateDisplay from './DateDisplay/DateDisplay';
import styles from './Header.module.css';
import UserInfo from './UserInfo/UserInfo';

function Header(): JSX.Element {
  return (
    <header className={styles.header}>
      <div>
        <h1 className={styles.logo}>Dashboard</h1>
      </div>
      <DateDisplay />
      <UserInfo />
    </header>
  );
}

export default Header;

import React from 'react';
import styles from './DateDisplay.module.css';

function DateDisplay(): JSX.Element {
  return (
    <div className={styles.dateField}>
      {formatDate(new Date())}
    </div>
  );
}

function formatDate(date: Date): string {
  const weekDays = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
  const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

  const day = date.getDate() <= 9
    ? '0' + date.getDate().toString()
    : date.getDate().toString();
  const weekDay = weekDays[date.getDay()];
  const month = months[date.getMonth()];

  return `${weekDay} ${day} ${month}`;
}

export default DateDisplay;

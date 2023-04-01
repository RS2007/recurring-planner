import dayjs from "dayjs";

/**
 *
 * @param day {number} day of the month
 * @param month  {number} month of the year
 * @param year  {number} year
 * @returns {string} A date string with given day,month and year
 */
export const dateStringConstructor = (day, month, year) => {
	return dayjs(`${year}-${month}-${day}`).format("YYYY-MM-DD");
};

export const numDaysInCurrentMonth = () =>
	new Date(new Date().getFullYear(), new Date().getMonth() + 1, 0).getDate();

export const getCurrentMonth = () => new Date().getMonth();

export const getCurrentYear = () => new Date().getFullYear();

export const getRange = (min: number, max: number) => {
	return [...Array(max - min + 1).keys()].map((i) => i + min);
};

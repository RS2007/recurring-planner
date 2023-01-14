import { user } from "../store/userStore";
import { accentColors } from "./formUtils";
type userCtx = {
  email: string;
  name: string;
  accentColor: "Blue" | "Pink";
};

export const getUser = (): userCtx => {
  let $user: userCtx;
  user.subscribe(($) => ($user = $ as userCtx))();
  return $user;
};

export const themeColor = () =>
  accentColors.find(
    (accentColor) => accentColor.label === getUser().accentColor
  ).class;

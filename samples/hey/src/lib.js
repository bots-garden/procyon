import { log } from "@suborbital/runnable";

export const run = (input) => {
  let message = "🎉 Hey (from JavaScript), " + input;

  log.info(message);

  return message;
};

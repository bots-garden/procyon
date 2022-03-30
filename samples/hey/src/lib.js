import { log } from "@suborbital/runnable";

export const run = (input) => {
  let message = "🎉 Hey, " + input;

  log.info(message);

  return message;
};

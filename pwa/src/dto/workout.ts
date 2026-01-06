import { Type } from "typebox";
import type { StaticParse } from "typebox";

type Workout = StaticParse<typeof Workout>;
const Workout = Type.Object({
    id: Type.String({ format: "uuid" }),
    name: Type.String(),
});

export { Workout };

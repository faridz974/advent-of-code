(() => {
  const solution = (right: number, bottom: number) : number => {
    const content = Deno.readTextFileSync("input.txt");
    const lines = content.split("\n");
    const map = new Array<Array<string>>();
    for (const l of lines.filter((l) => l !== "")) {
      const lineMap = new Array<string>();
      for (const char of l) {
        lineMap.push(char);
      }

      map.push(lineMap);
    }

    console.table(map);

    // start here
    let xIndex = 0;
    let yIndex = 0;
    let treeCount = 0;
    while (yIndex < map.length) {
      const currentPosition = map[yIndex][xIndex];
      if (currentPosition === "#") {
        treeCount++;
        console.log(treeCount);
      }

      xIndex = (xIndex + right) % map[yIndex].length;
      yIndex = yIndex + bottom;
    }

    return treeCount;
  };

  if (Deno.args[0] === "1") {
    console.log(solution(1, 2));
  } else if (Deno.args[0] === "2") {
    console.log(
        solution(1, 1) *
        solution(3, 1) *
        solution(5, 1) *
        solution(7, 1) *
        solution(1, 2));
  }

  Deno.exit(0);
})();

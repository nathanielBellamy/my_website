<script lang="ts">
  import { onMount } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"
  

  let x: number = 0
  let y: number = 0

  let state = rust.AppState.new()

  $: state_x = state.point.x
  $: state_y = state.point.y

  const handleClick = (e: any) => {
    x = e.clientX
    y = e.clientY - 143

    state.set_point(x, y)
    const newPointsLength = state.add_point(x,y)
    console.dir(newPoint)
  }

  let magicSquareWidth: number = 0
  let magicSquareHeight: number = 0

  onMount(async () => {
    let magicSquare = document.getElementById("magic_square_canvas_container")
    magicSquareWidth = magicSquare.offsetWidth
    magicSquareHeight = magicSquare.offsetHeight
	})

</script>

<div class="magic_square_container rounded-md flex flex-col justify-start">
  <div>
    {rust.AppState.foo("magicSquare bar")}
  </div>
  <div class="flex flex-row justify-around">
    <div>
     x:: {x}
    </div>
    <div>
     y:: {y}
    </div>
  </div>
  <div  class="magic_square_canvas_container grow"
        id="magic_square_canvas_container">
    <canvas id="magic_square"
            class="magic_square_canvas"
            width={magicSquareWidth} 
            height={magicSquareHeight}
            on:click={handleClick}/>
  </div>
</div>

<style lang="sass">
  .magic_square
    background-color: black

    &_container
      width: 100%
      height: 100%
      border: 2px solid black

    &_canvas
      background-color: gold
      width: 100%
      height: 100%
    
      &_container
        width: 100%
        background-color: red

</style>

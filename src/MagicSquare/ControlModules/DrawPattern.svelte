<script lang="ts">
  import { DrawPattern } from './DrawPattern'
  const drawPatternFormId: string = 'draw_pattern_form'
  const drawPatternHiddenInputId: string = 'magic_square_input_draw_pattern'

    // DRAW PATTERN
  // drawPattern vars
  enum DrawPatternDirection {
    Fix = "Fix",
    In = "In",
    Out = "Out"
  }

  export let drawPatternDirection: DrawPatternDirection
  export let drawPatternCount: number

  function deriveCurrDrawPattern(): DrawPattern {
    switch (drawPatternDirection) {
      case DrawPatternDirection.Fix:
        switch (drawPatternCount) {
          case 1:
            return DrawPattern.fix1
          case 2:
            return DrawPattern.fix2
          case 3:
            return DrawPattern.fix3
          case 4:
            return DrawPattern.fix4
          case 5:
            return DrawPattern.fix5
          case 6:
            return DrawPattern.fix6
          case 7:
            return DrawPattern.fix7
          case 8:
            return DrawPattern.fix8
        }
        break
      case DrawPatternDirection.In:
        switch (drawPatternCount) {
          case 1:
            return DrawPattern.in1
          case 2:
            return DrawPattern.in2
          case 3:
            return DrawPattern.in3
          case 4:
            return DrawPattern.in4
          case 5:
            return DrawPattern.in5
          case 6:
            return DrawPattern.in6
          case 7:
            return DrawPattern.in7
          case 8:
            return DrawPattern.in8
        }
        break
      case DrawPatternDirection.Out:
        switch (drawPatternCount) {
          case 1:
            return DrawPattern.out1
          case 2:
            return DrawPattern.out2
          case 3:
            return DrawPattern.out3
          case 4:
            return DrawPattern.out4
          case 5:
            return DrawPattern.out5
          case 6:
            return DrawPattern.out6
          case 7:
            return DrawPattern.out7
          case 8:
            return DrawPattern.out8
        }
        break
    }
    // default
    return DrawPattern.fix8
  }

  function setCurrDrawPatternDirection(direction: DrawPatternDirection) {
    drawPatternDirection = direction
  }

  function handleDrawPatternDirectionClick(direction: DrawPatternDirection) {
    setCurrDrawPatternDirection(direction)
    let form = document.getElementById(drawPatternFormId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  }

  function handleDrawPatternDirectionKeydown(e: any, direction: DrawPatternDirection) {
    if (e.keyCode === 13){
      setCurrDrawPatternDirection(direction)
      let form = document.getElementById(drawPatternFormId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  function setCurrDrawPatternCount(count: number) {
    drawPatternCount = count
  }

  function handleDrawPatternCountClick(count: number) {
    setCurrDrawPatternCount(count)
    let form = document.getElementById(drawPatternFormId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  }

  function handleDrawPatternCountKeydown(e: any, count: number) {
    if (e.keyCode === 13){
      setCurrDrawPatternCount(count)
      let form = document.getElementById(drawPatternFormId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  function handleDrawPatternFormSubmit(e: any) {
    e.preventDefault()
    var input = document.getElementById(drawPatternHiddenInputId)
    input.value = deriveCurrDrawPattern()
    input.dispatchEvent(new Event('input', {bubbles: true}))
    return false // do not refresh page
  }
</script>

<section class="flex flex-col justify-around items-stretch h-full w-full">
  <form id={drawPatternFormId}
        on:submit={handleDrawPatternFormSubmit}
        class="flex flex-col justify-around items-stretch h-full w-full">
    <div id="draw_pattern_buttons"
         class="grow flex flex-col justify-around items-stretch">
      <div id="draw_pattern_directions_outer"
           class="grow flex flex-col justify-around items-streth">
        <div id="draw_pattern_directions_inner"
             class="grow max-h-20 flex justify-around items-stretch">
          {#each Object.values(DrawPatternDirection) as dir}
            <button class="grow max-h-26 pr-3 pl-3"
                    on:click={() => handleDrawPatternDirectionClick(dir)}
                    on:keydown={(e) => handleDrawPatternDirectionKeydown(e, dir)}
                    class:selected={drawPatternDirection === dir}>
              {dir}
            </button>
          {/each}
        </div>
      </div>
      <div id="draw_pattern_counts"
           class="grow flex flex-col justify-around items-stretch">
        {#each [0,4] as countShifter}
          <div id="draw_pattern_counts_row"
               class="grow flex justify-evenly items-stretch gap-0">
            {#each [1,2,3,4].map(x => x + countShifter) as count}
              <button class="grow max-h-20"
                      on:click={() => handleDrawPatternCountClick(count)}
                      on:keydown={(e) => handleDrawPatternCountKeydown(e, count)}
                      class:selected={drawPatternCount === count}>
                {count}
              </button>
            {/each}
          </div>
        {/each}
      </div>
    </div>
    <slot name="hiddenInput" />
  </form>
</section>

<style lang="sass">
  @use "../../styles/color"
  .selected
    background-color: color.$blue-8
</style>

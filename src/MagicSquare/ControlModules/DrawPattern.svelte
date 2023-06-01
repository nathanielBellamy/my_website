<script lang="ts">
  const drawPatternFormId: string = 'draw_pattern_form'
  const drawPatternHiddenInputId: string = 'magic_square_input_draw_pattern'

    // DRAW PATTERN
  // drawPattern vars
  enum DrawPatternDirection {
    Fix = "Fix",
    In = "In",
    Out = "Out"
  }

  export let currDrawPatternDirection: DrawPatternDirection
  export let currDrawPatternCount: number

  function deriveCurrDrawPattern(): string {
    var result: string
    switch (currDrawPatternDirection) {
      case DrawPatternDirection.Fix:
        result = DrawPatternDirection.Fix
        break
      case DrawPatternDirection.In:
        result = DrawPatternDirection.In
        break
      case DrawPatternDirection.Out:
        result = DrawPatternDirection.Out
        break
    }

    return `${result}${currDrawPatternCount}`
  }

  function setCurrDrawPatternDirection(direction: DrawPatternDirection) {
    currDrawPatternDirection = direction
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
    currDrawPatternCount = count
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

  function handleDrawPatternFormSubmit() {
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
         class="h-full flex flex-col justify-around">
      <div id="draw_pattern_directions_outer"
           class="grow flex flex-col justify-around items-streth">
        <div id="draw_pattern_directions_inner"
             class="grow max-h-20 flex justify-around items-stretch">
          {#each Object.values(DrawPatternDirection) as dir}
            <button class="grow max-h-26 pr-3 pl-3"
                    on:click={() => handleDrawPatternDirectionClick(dir)}
                    on:keydown={(e) => handleDrawPatternDirectionKeydown(e, dir)}
                    class:selected={currDrawPatternDirection === dir}>
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
                      class:selected={currDrawPatternCount === count}>
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

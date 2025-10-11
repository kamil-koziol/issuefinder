<button {...$$props}>
  <slot />
</button>

<style>
  button {
    all: unset;
    color: white;
    background-color: var(--primary);
    padding: 8px 18px;
    border-radius: 8px;
    cursor: pointer;
    font-size: 14px;
  }
</style>

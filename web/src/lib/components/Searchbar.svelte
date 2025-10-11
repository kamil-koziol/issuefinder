<script>
  import SearchIcon from '$lib/assets/icons/SearchIcon.svelte';
</script>

<div class="searchHolder">
  <SearchIcon width="18px" title="Search" />
  <input type="text" id="searchbar" placeholder="Search issues..." />
</div>

<style>
  .searchHolder {
    display: flex;
    align-items: center;
    background-color: #f1f3f4;
    padding: 2px 8px;
    border-radius: 10px;
  }

  input {
    all: unset;
    min-width: 60vw;
    padding: 4px 12px;
    border: unset;
    background-color: #f1f3f4;
    font-size: 14px;
    color: #333;
  }
</style>

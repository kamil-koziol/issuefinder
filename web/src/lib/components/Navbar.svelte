<script>
  import Searchbar from '$lib/components/Searchbar.svelte';
  import Button from '$lib/components/Button.svelte';
  import Logo from '$lib/assets/icons/Logo.svelte';
</script>

<nav>
  <div class="leftSide">
    <Logo width="40px" />
    <h1>IssueFinder</h1>
  </div>
  <Searchbar />
  <Button>Sign in</Button>
</nav>

<style>
  nav {
    display: flex;
    align-items: center;
    justify-content: space-between;
    box-shadow:
      0 1px 2px 0 hsla(0, 0%, 50%, 0.3),
      0 2px 6px 2px hsla(0, 0%, 50%, 0.15);
    padding: 0px 10px;
  }

  h1 {
    font-size: 20px;
    font-weight: normal;
    color: #333;
  }

  .leftSide {
    display: flex;
    gap: 10px;
    align-items: center;
  }
</style>

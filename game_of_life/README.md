# Game of Life

<p><strong>About this Kata</strong></p>

<p>This was presented as a <a href="https://codingdojo.org/PreparedKata">PreparedKata</a> at XP2005, by Emmanuel Gaillot.</p>

<p>Difficulty - medium</p>

<p>Similar Katas - <a href="https://codingdojo.org/kata/Minesweeper">KataMinesweeper</a> ,
<a href="https://codingdojo.org/kata/Reversi">KataReversi</a></p>

<h1 id="problem-description">Problem Description</h1>

<p>This Kata is about calculating the next generation of Conway’s game of
life, given any starting position. See
<a href="http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life">http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life</a> for background.</p>

<p>You start with a two dimensional grid of cells, where each cell is
either alive or dead. In this version of the problem, the grid is
finite, and no life can exist off the edges. When calcuating the next
generation of the grid, follow these rules:</p>

<pre><code>   1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
   2. Any live cell with more than three live neighbours dies, as if by overcrowding.
   3. Any live cell with two or three live neighbours lives on to the next generation.
   4. Any dead cell with exactly three live neighbours becomes a live cell.
</code></pre>

<p>You should write a program that can accept an arbitrary grid of cells,
and will output a similar grid showing the next generation.</p>

<h1 id="clues">Clues</h1>

<p>The input starting position could be a text file that looks like this:</p>

<pre><code>Generation 1:
4 8
........
....*...
...**...
........
</code></pre>

<p>And the output could look like this:</p>

<pre><code>Generation 2:
4 8
........
...**...
...**...
........
</code></pre>

<p>The input format is similar to that in <a href="https://codingdojo.org/kata/Minesweeper">KataMinesweeper</a>, and is easily parsed in most languages.</p>

<h1 id="front-end">Front End</h1>

<p>Make a beautiful IHM to render grid of any size in one page.</p>

<h1 id="suggested-test-cases">Suggested Test Cases</h1>

<p>Make sure you have enough coverage of edge</p>

<p>cases - where there are births and deaths at the edge of the grid.</p>

<h1 id="comments-from-those-who-are-working-on-this-kata">Comments from those who are working on this Kata</h1>

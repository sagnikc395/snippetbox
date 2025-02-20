import { Hono } from 'hono';

const app = new Hono();

//handle POST, PUT ,DELETE also

app.post('/posts', (c) => c.text('Created!', 201));

app.delete('/posts/:id', (c) => {
	return c.text(`${c.req.param('id')} is deleted!`);
});

app.get('/api/hello', (c) => {
	return c.json({
		ok: true,
		message: 'Hello Hono!'
	});
});

app.get('/', (c) => {
	return c.text('Hello Hono!');
});

export default app;

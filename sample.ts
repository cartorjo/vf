// sample.ts

function greet(name: string): string {
    return `Hello, ${name}!`;
}

const personName: string = "Jose";
const greeting: string = greet(personName);

console.log(greeting);

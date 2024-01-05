const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

const answer = () =>
    new Promise((resolve) =>
        rl.question('Enter password:', (reply) => resolve(reply))
    );

class WorkSecurityPass {
    date = new Date();

    constructor(name, role, accessKey) {
        this.name = name;
        this.role = role;
        this.accessKey = accessKey;
        this.registeredDate = this.date.toLocaleString();
        this.lastLogin = null;
        this.lastLogout = null;
        this.timeline = [];
    }

    async loginAccess(password) {
        const date = new Date();
        let verifyPass = true;
        while (verifyPass) {
            const response = await answer();
            if (response == this.accessKey) {
                verifyPass = false;
            } else {
                console.log('Incorrect password. Try again');
            }
        }
        rl.close();

        if (date.toLocaleDateString() == this.lastLogin?.toLocaleDateString()) {
            console.log("You're logged in already!");
            return;
        }
        const login = {};
        const day = date.toDateString();
        login[day] = {};
        login[day]['resumed'] = date.toLocaleTimeString();
        this.lastLogin = date.toLocaleString();
        this.timeline.push(login);
        console.log('Entry access granted');
        return;
    }

    logoutTime() {
        // do some other oprations
    }
}

const newStaff = new WorkSecurityPass('toby', 'CEO', 'tobytoby');
newStaff.loginAccess();
// await taptap();
// console.log('newStaff data', newStaff.timeline);

async function taptap() {
    return;
}

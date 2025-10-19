1. Analysis document

Έχουμε στο repository 3 έγγραφα:
- Το ένα είναι το go πρόγραμμα
- Το δεύτερο είναι το README.md έγγραφο που αναλύει τι κάνει το πρόγραμμα και ποιες
είναι οι συνθήκες του
- Ένα example.txt έγγραφο για να βοηθήσει τον χρήστη να καταλάβει καλύτερα

Ο χρήστης που θέλει να χρησιμοποιήσει το πρόγραμμα θα χρειαστεί να
δημιουργήσει/μεταφέρει στο ίδιο directory ένα
έγγραφο .txt το οποίο θα μπορεί να διαβάζει τα modifiers μέσα στις παρενθέσεις και θα
κάνει τις απαραίτητες αλλαγές.

- Άμα μέσα στην παρένθεση υπάρχει η λέξη 'hex' τότε θα μετατρέπει την αμέσως
προηγούμενη λέξη στην δεκαδική μορφή του (η λέξη αυτή θα είναι πάντα ένας δεκαεξαδικός
αριθμός ώστε να είναι εφικτή η μετατροπή του)
π.χ. "1E (hex) files were added" -> "30 files were added"

- Άμα μέσα στην παρένθεση υπάρχει η λέξη 'bin' τότε θα μετατρέπει την αμέσως
προηγούμενη λέξη στην δεκαδική μορφή του (η λέξη αυτή θα είναι πάντα ένας δυαδικός αριθμός ώστε να είναι εφικτή η μετατροπή του)
π.χ. "It has been 10 (bin) years" -> "It has been 2 years"

- Άμα μέσα στην παρένθεση υπάρχει η λέξη 'up' τότε θα μετατρέπει όλα τα γράμματα της
αμέσως προηγούμενης λέξης σε κεφαλαία
π.χ."Ready, set, go (up) !" -> "Ready, set, GO!"

- Άμα μέσα στην παρένθεση υπάρχει η λέξη 'low' τότε θα μετατρέπει όλα τα γράμματα της
αμέσως προηγούμενης λέξης που είναι κεφαλαία σε μικρά
π.χ. "I should stop SHOUTING (low)" -> "I should stop shouting"

- Άμα μέσα στην παρένθεση υπάρχει η λέξη 'cap' τότε θα μετατρέπει το πρώτο γράμμα της
αμέσως προηγούμενης λέξης σε κεφαλαίο
π.χ. "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"

** Σε περίπτωση που μέσα στις παρενθέσεις βρίσκονται και αριθμοί σε
μορφή (<λέξη>, <αριθμός>), τότε η μετατροπή ισχύει για τις <αριθμός> λέξεις
πριν την παρένθεση
π.χ. "This is so exciting (up, 2)" -> "This is SO EXCITING"

- Κάθε ",", ".", "!", "?", ":" και ";" πρέπει να είναι κοντά στην προηγούμενη λέξη και
να απέχει ένα κενό με την επόμενη
π.χ. "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"

** Σε περίπτωση που υπάρχει "..." ή "!?" θα ομαδοποιούνται και θα ισχύει ο παραπάνω
κανόνας κανονικά
π.χ. "I was thinking ... You were right" -> "I was thinking... You were right"

- Κάθε "'" θα είναι μαζί με ένα δεύτερο "'" και θα πρέπει να βρίσκονται δεξιά και αριστερά
της λέξης που περιέχουν, χωρίς κανένα κενό ανάμεσα
π.χ. "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'"

** Αν υπάρχει περισσότερη από μία λέξη μέσα στα ' ' τότε θα πρέπει να μην υπάρχει κανένα
κενό ανάμεσα στα "'" και στη λέξη δίπλα του
π.χ. "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'"

- Τέλος, θα μετατρέπει το 'a' σε 'an' αν η επόμενη λέξη ξεκινάει με φωνήεν.
π.χ. "There it was. A amazing rock!" -> "There it was. An amazing rock!"

----------- Διαφορά pipeline με FSM --------------

Οι αρχιτεκτονικές Pipeline και FSM (Finite State Machine) είναι δύο τρόποι να φτιάξεις ένα σύστημα που κάνει δουλειές βήμα-βήμα.

- Pipeline: Χωρίζει τη δουλειά σε στάδια, και κάθε στάδιο δουλεύει ταυτόχρονα με τα άλλα — σαν εργοστάσιο όπου κάθε εργάτης κάνει ένα μέρος της δουλειάς.
Έτσι, γίνεται πιο γρήγορη η συνολική εργασία, αλλά χρησιμοποιείται πιο πολύ μνήμη

- FSM: Το σύστημα κάνει ένα βήμα τη φορά, αλλάζει «κατάσταση» ανάλογα με το τι συμβαίνει. Είναι συνήθως πιο αργό, αλλά πολύ αποτελεσματικό αν θέλουμε να
χρησιμοποιήσουμε όσο πιο λίγη μνήμη γίνεται

Προσωπική επιλογή: Θα διάλεγα το Pipeline, γιατί επιτρέπει να γίνονται πολλά πράγματα ταυτόχρονα, κάνοντας το σύστημα πιο γρήγορο και αποδοτικό, το οποίο είναι
σημαντικότερο για μένα από την αποδοτικότητα της μνήμης

---------------------------------------------------------------------------------------------------------------------------------------------------------------------

2. "Golden Test Set" (Success Test Cases)

Βασικά test cases από τα audit examples του project:

1) If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?

Σκοπός είναι να ελεγχθεί ότι το σύστημα εφαρμόζει σωστά τις εντολές αλλαγής μορφοποίησης (κεφαλαία/πεζά)
στις λέξεις μιας πρότασης.

Το σύστημα εντοπίζει την εντολή:
i) (low, 3) και μετατρέπει τις τρεις λέξεις που προηγούνται (BREAKFAST, IN, BED) σε πεζά -> breakfast, in, bed
ii) (cap) και μετατρέπει το πρώτο γράμμα της λέξης how σε κεφαλαίο -> How.
iii) (up, 2) και μετατρέπει τις δύο επόμενες λέξεις (my, house) σε κεφαλαία -> MY HOUSE.

Άρα το αποτέλεσμα θα πρέπει να είναι:
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?

2) I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure.

Σκοπός είναι να ελεγχθεί ότι το σύστημα εφαρμόζει σωστά τις εντολές μετατροπής
αριθμητικών μορφών (δυαδικό, δεκαεξαδικό) σε δεκαδική μορφή μέσα σε μια πρόταση.

Το σύστημα εντοπίζει την εντολή:
i) (bin) και μετατρέπει τον δυαδικό αριθμό 101 σε δεκαδικό → 5.
ii) (hex) και μετατρέπει τον δεκαεξαδικό αριθμό 1a σε δεκαδικό → 26.

Άρα το αποτέλεσμα θα πρέπει να είναι:
I have to pack 5 outfits. Packed 26 just to be sure.

3) Don not be sad ,because sad backwards is das . And das not good

Σκοπός είναι να ελεγχθεί ότι το σύστημα εφαρμόζει σωστά τη διόρθωση στίξης μέσα
στην πρόταση (διαχείριση κενών πριν και μετά από σημεία στίξης).

Το σύστημα εντοπίζει τα λάθη στίξης και:
i) Αφαιρεί το περιττό κενό πριν από το κόμμα μετά τη λέξη sad → sad, because...
ii) Αφαιρεί το περιττό κενό πριν από την τελεία μετά τη λέξη das → das.
** Μετά από κάθε σημείο στίξης θα πρέπει να υπάρχει ένα κενό αν υπάρχει και άλλη λέξη έπειτα**

Άρα το αποτέλεσμα θα πρέπει να είναι:
Don not be sad, because sad backwards is das. And das not good

4) harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
Σκοπός είναι να ελεγχθεί ότι το σύστημα εφαρμόζει σωστά τις εντολές μορφοποίησης κειμένου
και στίξης (κεφαλαίο πρώτο γράμμα, άρθρα, κόμματα, τελείες) μέσα στην πρόταση.

Το σύστημα εντοπίζει τις εντολές:
i) (cap, 2) και μετατρέπει τις δύο πρώτες λέξεις harold wilson ώστε να αρχίζουν με κεφαλαίο → Harold Wilson.
ii) Αφαιρεί τα περιττά κενά μετά τα σημεία στίξης (π.χ. μετά το κόμμα και πριν από την τελεία).
iii) Διορθώνει το άρθρο a σε an πριν από λέξη που αρχίζει με φωνήεν → an optimist.
iv) Τοποθετεί σωστά τα σημεία στίξης και αφαιρεί τα περιττά κενά μέσα στα εισαγωγικά

Άρα το αποτέλεσμα θα πρέπει να είναι:
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'

** Παραδείγματα tricky περιπτώσεων που θα χρειαστεί να αντιμετωπίσουμε:

- i love PYtHOn (cap)!

Εδώ θα χρειαστεί να μετατρέψουμε όλα τα γράμματα που είναι κεφαλαία (εκτός από την πρώτη)
σε πεζά, καθώς το cap μετατρέπει ολόκληρη την λέξη σε capitalized μορφή

- the playground was too easy ( up)

Εδώ δεν θα πρέπει να γίνει καμία αλλαγή, καθώς ο χρήστης δεν έγραψε σωστά το format του modifier

- i love golang (up, 10)

Θα πρέπει ο κώδικας να μην εμφανίζει κάποιο error όταν ο χρήστης θέλει να κάνει modify
περισσότερες από τις υπάρχουσες λέξεις

- i love typescript (up, 0)

Δεν θα πρέπει να γίνεται καμία αλλαγή

- i love javascript (up, -1) a lot

Εδώ θα πρέπει η επόμενη λέξη να αλλάζει, δηλαδή το a, όχι η προηγούμενη

-------------------------------------------------------------------

Μία παράγραφος για το sample.txt

The system received 2F (hex) REQUESTS fROM mUlTiPle USERS (cap, 4) , but
1101 (bin) of them failed (up, 3) ! Meanwhile, the admin noticed that a
unusual alert (low, 2) had triggered on the main server . He immediately
said: ' a critical error has occurred ! ' (cap, 5) ... The security team
analyzed 3C (hex) logs and discovered that 101 (bin) attempts were made to access
restricted files ; surprisingly, some files contained ' confidential information ' (low, 4) .
After reviewing the logs, they realized that the failed attempts happened between 12:00 PM and 3:00 PM ,
and they decided to notify all affected users . One user exclaimed: ' a unbelievable coincidence ! ' , while another
added: ' i had no idea this could happen ! ' The system then generated 2 (bin) alerts for minor warnings , and a summary
report was emailed to the management team . In total, 4A (hex) records were successfully processed , and the team agreed
that ' proper monitoring procedures ' (cap, -3) should be enforced strictly in the future .

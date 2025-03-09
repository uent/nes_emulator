| Opcode | Mnemonic | Addressing       | Bytes | Cycles                                    | Descripción                                                         |
|--------|----------|------------------|-------|-------------------------------------------|---------------------------------------------------------------------|
| $00    | BRK      | Implied          | 1     | 7                                         | Fuerza una interrupción software.                                   |
| $01    | ORA      | (Indirect,X)     | 2     | 6                                         | OR lógico entre el acumulador y la memoria (indirección pre-indexada).|
| $02    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $03    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $04    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $05    | ORA      | Zero Page        | 2     | 3                                         | OR lógico en dirección de cero página.                             |
| $06    | ASL      | Zero Page        | 2     | 5                                         | Desplaza a la izquierda el contenido en cero página.                |
| $07    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $08    | PHP      | Implied          | 1     | 3                                         | Empuja el registro de estado a la pila.                             |
| $09    | ORA      | Immediate        | 2     | 2                                         | OR lógico entre el acumulador y un valor inmediato.                 |
| $0A    | ASL      | Accumulator      | 1     | 2                                         | Desplaza a la izquierda el acumulador.                              |
| $0B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $0C    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $0D    | ORA      | Absolute         | 3     | 4                                         | OR lógico en dirección absoluta.                                    |
| $0E    | ASL      | Absolute         | 3     | 6                                         | Desplaza a la izquierda el valor en dirección absoluta.             |
| $0F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $10    | BPL      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag negativo es 0 (bifurcación).                         |
| $11    | ORA      | (Indirect),Y    | 2     | 5+ if page crossed                         | OR lógico con indirección post-indexada en Y.                        |
| $12    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $13    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $14    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $15    | ORA      | Zero Page,X      | 2     | 4                                         | OR lógico en cero página con índice X.                              |
| $16    | ASL      | Zero Page,X      | 2     | 6                                         | Desplaza a la izquierda en cero página con índice X.                |
| $17    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $18    | CLC      | Implied          | 1     | 2                                         | Borra el flag de acarreo.                                             |
| $19    | ORA      | Absolute,Y       | 3     | 4+ if page crossed                         | OR lógico en dirección absoluta con índice Y.                       |
| $1A    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $1B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $1C    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $1D    | ORA      | Absolute,X       | 3     | 4+ if page crossed                         | OR lógico en dirección absoluta con índice X.                       |
| $1E    | ASL      | Absolute,X       | 3     | 7                                         | Desplaza a la izquierda en dirección absoluta con índice X.         |
| $1F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $20    | JSR      | Absolute         | 3     | 6                                         | Llama a una subrutina, guardando la dirección de retorno.             |
| $21    | AND      | (Indirect,X)     | 2     | 6                                         | AND lógico entre el acumulador y la memoria (indirección pre-indexada).|
| $22    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $23    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $24    | BIT      | Zero Page        | 2     | 3                                         | Prueba de bits: afecta flags sin modificar el acumulador.           |
| $25    | AND      | Zero Page        | 2     | 3                                         | AND lógico en cero página.                                           |
| $26    | ROL      | Zero Page        | 2     | 5                                         | Rota a la izquierda el valor en cero página a través del acarreo.     |
| $27    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $28    | PLP      | Implied          | 1     | 4                                         | Recupera el registro de estado desde la pila.                        |
| $29    | AND      | Immediate        | 2     | 2                                         | AND lógico entre el acumulador y un valor inmediato.                |
| $2A    | ROL      | Accumulator      | 1     | 2                                         | Rota a la izquierda el acumulador a través del acarreo.               |
| $2B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $2C    | BIT      | Absolute         | 3     | 4                                         | Prueba de bits en dirección absoluta.                               |
| $2D    | AND      | Absolute         | 3     | 4                                         | AND lógico en dirección absoluta.                                  |
| $2E    | ROL      | Absolute         | 3     | 6                                         | Rota a la izquierda el valor en dirección absoluta.                 |
| $2F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $30    | BMI      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag negativo es 1 (bifurcación).                         |
| $31    | AND      | (Indirect),Y    | 2     | 5+ if page crossed                         | AND lógico con indirección post-indexada en Y.                       |
| $32    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $33    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $34    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $35    | AND      | Zero Page,X      | 2     | 4                                         | AND lógico en cero página con índice X.                             |
| $36    | ROL      | Zero Page,X      | 2     | 6                                         | Rota a la izquierda en cero página con índice X.                    |
| $37    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $38    | SEC      | Implied          | 1     | 2                                         | Establece el flag de acarreo.                                         |
| $39    | AND      | Absolute,Y       | 3     | 4+ if page crossed                         | AND lógico en dirección absoluta con índice Y.                      |
| $3A    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $3B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $3C    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $3D    | AND      | Absolute,X       | 3     | 4+ if page crossed                         | AND lógico en dirección absoluta con índice X.                      |
| $3E    | ROL      | Absolute,X       | 3     | 7                                         | Rota a la izquierda en dirección absoluta con índice X.             |
| $3F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $40    | RTI      | Implied          | 1     | 6                                         | Retorna de una interrupción, restaurando el estado.                 |
| $41    | EOR      | (Indirect,X)     | 2     | 6                                         | XOR exclusivo entre el acumulador y la memoria (indirección pre-indexada).|
| $42    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $43    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $44    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $45    | EOR      | Zero Page        | 2     | 3                                         | XOR exclusivo en cero página.                                       |
| $46    | LSR      | Zero Page        | 2     | 5                                         | Desplaza a la derecha el valor en cero página.                      |
| $47    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $48    | PHA      | Implied          | 1     | 3                                         | Empuja el acumulador a la pila.                                     |
| $49    | EOR      | Immediate        | 2     | 2                                         | XOR exclusivo entre el acumulador y un valor inmediato.             |
| $4A    | LSR      | Accumulator      | 1     | 2                                         | Desplaza a la derecha el acumulador.                                |
| $4B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $4C    | JMP      | Absolute         | 3     | 3                                         | Salta a la dirección absoluta especificada.                         |
| $4D    | EOR      | Absolute         | 3     | 4                                         | XOR exclusivo en dirección absoluta.                                |
| $4E    | LSR      | Absolute         | 3     | 6                                         | Desplaza a la derecha el valor en dirección absoluta.               |
| $4F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $50    | BVC      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag de desbordamiento es 0 (bifurcación).                |
| $51    | EOR      | (Indirect),Y    | 2     | 5+ if page crossed                         | XOR exclusivo usando indirección post-indexada en Y.                 |
| $52    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $53    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $54    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $55    | EOR      | Zero Page,X      | 2     | 4                                         | XOR exclusivo en cero página con índice X.                          |
| $56    | LSR      | Zero Page,X      | 2     | 6                                         | Desplaza a la derecha en cero página con índice X.                  |
| $57    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $58    | CLI      | Implied          | 1     | 2                                         | Borra el flag de interrupción.                                        |
| $59    | EOR      | Absolute,Y       | 3     | 4+ if page crossed                         | XOR exclusivo en dirección absoluta con índice Y.                   |
| $5A    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $5B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $5C    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $5D    | EOR      | Absolute,X       | 3     | 4+ if page crossed                         | XOR exclusivo en dirección absoluta con índice X.                   |
| $5E    | LSR      | Absolute,X       | 3     | 7                                         | Desplaza a la derecha en dirección absoluta con índice X.           |
| $5F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $60    | RTS      | Implied          | 1     | 6                                         | Retorna de una subrutina.                                             |
| $61    | ADC      | (Indirect,X)     | 2     | 6                                         | Suma con acarreo usando indirección pre-indexada.                    |
| $62    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $63    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $64    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $65    | ADC      | Zero Page        | 2     | 3                                         | Suma con acarreo en cero página.                                    |
| $66    | ROR      | Zero Page        | 2     | 5                                         | Rota a la derecha el valor en cero página a través del acarreo.       |
| $67    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $68    | PLA      | Implied          | 1     | 4                                         | Recupera el acumulador de la pila.                                  |
| $69    | ADC      | Immediate        | 2     | 2                                         | Suma con acarreo con un valor inmediato.                            |
| $6A    | ROR      | Accumulator      | 1     | 2                                         | Rota a la derecha el acumulador.                                     |
| $6B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $6C    | JMP      | Indirect         | 3     | 5                                         | Salta a la dirección indirecta.                                     |
| $6D    | ADC      | Absolute         | 3     | 4                                         | Suma con acarreo en dirección absoluta.                             |
| $6E    | ROR      | Absolute         | 3     | 6                                         | Rota a la derecha el valor en dirección absoluta.                   |
| $6F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $70    | BVS      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag de desbordamiento es 1 (bifurcación).                |
| $71    | ADC      | (Indirect),Y    | 2     | 5+ if page crossed                         | Suma con acarreo usando indirección post-indexada en Y.              |
| $72    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $73    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $74    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $75    | ADC      | Zero Page,X      | 2     | 4                                         | Suma con acarreo en cero página con índice X.                        |
| $76    | ROR      | Zero Page,X      | 2     | 6                                         | Rota a la derecha en cero página con índice X.                       |
| $77    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $78    | SEI      | Implied          | 1     | 2                                         | Establece el flag de interrupción.                                   |
| $79    | ADC      | Absolute,Y       | 3     | 4+ if page crossed                         | Suma con acarreo en dirección absoluta con índice Y.                 |
| $7A    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $7B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $7C    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $7D    | ADC      | Absolute,X       | 3     | 4+ if page crossed                         | Suma con acarreo en dirección absoluta con índice X.                 |
| $7E    | ROR      | Absolute,X       | 3     | 7                                         | Rota a la derecha en dirección absoluta con índice X.                |
| $7F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $80    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $81    | STA      | (Indirect,X)     | 2     | 6                                         | Almacena el acumulador usando indirección pre-indexada.               |
| $82    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $83    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $84    | STY      | Zero Page        | 2     | 3                                         | Almacena el registro Y en cero página.                               |
| $85    | STA      | Zero Page        | 2     | 3                                         | Almacena el acumulador en cero página.                               |
| $86    | STX      | Zero Page        | 2     | 3                                         | Almacena el registro X en cero página.                               |
| $87    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $88    | DEY      | Implied          | 1     | 2                                         | Decrementa el registro Y.                                             |
| $89    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $8A    | TXA      | Implied          | 1     | 2                                         | Transfiere el registro X al acumulador.                              |
| $8B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $8C    | STY      | Absolute         | 3     | 4                                         | Almacena el registro Y en dirección absoluta.                        |
| $8D    | STA      | Absolute         | 3     | 4                                         | Almacena el acumulador en dirección absoluta.                        |
| $8E    | STX      | Absolute         | 3     | 4                                         | Almacena el registro X en dirección absoluta.                        |
| $8F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $90    | BCC      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag de acarreo es 0 (bifurcación).                        |
| $91    | STA      | (Indirect),Y    | 2     | 6                                         | Almacena el acumulador usando indirección post-indexada en Y.          |
| $92    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $93    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $94    | STY      | Zero Page,X      | 2     | 4                                         | Almacena el registro Y en cero página con índice X.                  |
| $95    | STA      | Zero Page,X      | 2     | 4                                         | Almacena el acumulador en cero página con índice X.                  |
| $96    | STX      | Zero Page,Y      | 2     | 4                                         | Almacena el registro X en cero página con índice Y.                  |
| $97    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $98    | TYA      | Implied          | 1     | 2                                         | Transfiere el registro Y al acumulador.                              |
| $99    | STA      | Absolute,Y       | 3     | 5                                         | Almacena el acumulador en dirección absoluta con índice Y.           |
| $9A    | TXS      | Implied          | 1     | 2                                         | Transfiere el registro X al stack pointer.                           |
| $9B    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $9C    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $9D    | STA      | Absolute,X       | 3     | 5                                         | Almacena el acumulador en dirección absoluta con índice X.           |
| $9E    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $9F    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $A0    | LDY      | Immediate        | 2     | 2                                         | Carga un valor inmediato en el registro Y.                           |
| $A1    | LDA      | (Indirect,X)     | 2     | 6                                         | Carga el acumulador usando indirección pre-indexada.                 |
| $A2    | LDX      | Immediate        | 2     | 2                                         | Carga un valor inmediato en el registro X.                           |
| $A3    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $A4    | LDY      | Zero Page        | 2     | 3                                         | Carga el registro Y desde cero página.                               |
| $A5    | LDA      | Zero Page        | 2     | 3                                         | Carga el acumulador desde cero página.                               |
| $A6    | LDX      | Zero Page        | 2     | 3                                         | Carga el registro X desde cero página.                               |
| $A7    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $A8    | TAY      | Implied          | 1     | 2                                         | Transfiere el acumulador al registro Y.                              |
| $A9    | LDA      | Immediate        | 2     | 2                                         | Carga un valor inmediato en el acumulador.                           |
| $AA    | TAX      | Implied          | 1     | 2                                         | Transfiere el acumulador al registro X.                              |
| $AB    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $AC    | LDY      | Absolute         | 3     | 4                                         | Carga el registro Y desde dirección absoluta.                        |
| $AD    | LDA      | Absolute         | 3     | 4                                         | Carga el acumulador desde dirección absoluta.                        |
| $AE    | LDX      | Absolute         | 3     | 4                                         | Carga el registro X desde dirección absoluta.                        |
| $AF    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $B0    | BCS      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag de acarreo es 1 (bifurcación).                        |
| $B1    | LDA      | (Indirect),Y    | 2     | 5+ if page crossed                         | Carga el acumulador usando indirección post-indexada en Y.             |
| $B2    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $B3    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $B4    | LDY      | Zero Page,X      | 2     | 4                                         | Carga el registro Y desde cero página con índice X.                  |
| $B5    | LDA      | Zero Page,X      | 2     | 4                                         | Carga el acumulador desde cero página con índice X.                  |
| $B6    | LDX      | Zero Page,Y      | 2     | 4                                         | Carga el registro X desde cero página con índice Y.                  |
| $B7    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $B8    | CLV      | Implied          | 1     | 2                                         | Borra el flag de desbordamiento.                                      |
| $B9    | LDA      | Absolute,Y       | 3     | 4+ if page crossed                         | Carga el acumulador desde dirección absoluta con índice Y.           |
| $BA    | TSX      | Implied          | 1     | 2                                         | Transfiere el stack pointer al registro X.                           |
| $BB    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $BC    | LDY      | Absolute,X       | 3     | 4+ if page crossed                         | Carga el registro Y desde dirección absoluta con índice X.           |
| $BD    | LDA      | Absolute,X       | 3     | 4+ if page crossed                         | Carga el acumulador desde dirección absoluta con índice X.           |
| $BE    | LDX      | Absolute,Y       | 3     | 4+ if page crossed                         | Carga el registro X desde dirección absoluta con índice Y.           |
| $BF    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $C0    | CPY      | Immediate        | 2     | 2                                         | Compara el registro Y con un valor inmediato.                        |
| $C1    | CMP      | (Indirect,X)     | 2     | 6                                         | Compara el acumulador con la memoria (indirección pre-indexada).       |
| $C2    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $C3    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $C4    | CPY      | Zero Page        | 2     | 3                                         | Compara el registro Y con un valor en cero página.                    |
| $C5    | CMP      | Zero Page        | 2     | 3                                         | Compara el acumulador con un valor en cero página.                    |
| $C6    | DEC      | Zero Page        | 2     | 5                                         | Decrementa el valor en una dirección de cero página.                 |
| $C7    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $C8    | INY      | Implied          | 1     | 2                                         | Incrementa el registro Y.                                             |
| $C9    | CMP      | Immediate        | 2     | 2                                         | Compara el acumulador con un valor inmediato.                        |
| $CA    | DEX      | Implied          | 1     | 2                                         | Decrementa el registro X.                                             |
| $CB    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $CC    | CPY      | Absolute         | 3     | 4                                         | Compara el registro Y con un valor en dirección absoluta.             |
| $CD    | CMP      | Absolute         | 3     | 4                                         | Compara el acumulador con un valor en dirección absoluta.             |
| $CE    | DEC      | Absolute         | 3     | 6                                         | Decrementa el valor en dirección absoluta.                           |
| $CF    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $D0    | BNE      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag cero es 0 (bifurcación).                             |
| $D1    | CMP      | (Indirect),Y    | 2     | 5+ if page crossed                         | Compara el acumulador con la memoria usando indirección post-indexada en Y.|
| $D2    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $D3    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $D4    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $D5    | CMP      | Zero Page,X      | 2     | 4                                         | Compara el acumulador con un valor en cero página con índice X.        |
| $D6    | DEC      | Zero Page,X      | 2     | 6                                         | Decrementa el valor en cero página con índice X.                     |
| $D7    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $D8    | CLD      | Implied          | 1     | 2                                         | Borra el flag decimal.                                                |
| $D9    | CMP      | Absolute,Y       | 3     | 4+ if page crossed                         | Compara el acumulador con un valor en dirección absoluta con índice Y.  |
| $DA    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $DB    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $DC    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $DD    | CMP      | Absolute,X       | 3     | 4+ if page crossed                         | Compara el acumulador con un valor en dirección absoluta con índice X.  |
| $DE    | DEC      | Absolute,X       | 3     | 7                                         | Decrementa el valor en dirección absoluta con índice X.              |
| $DF    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $E0    | CPX      | Immediate        | 2     | 2                                         | Compara el registro X con un valor inmediato.                        |
| $E1    | SBC      | (Indirect,X)     | 2     | 6                                         | Resta con acarreo usando indirección pre-indexada.                    |
| $E2    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $E3    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $E4    | CPX      | Zero Page        | 2     | 3                                         | Compara el registro X con un valor en cero página.                    |
| $E5    | SBC      | Zero Page        | 2     | 3                                         | Resta con acarreo en cero página.                                     |
| $E6    | INC      | Zero Page        | 2     | 5                                         | Incrementa el valor en una dirección de cero página.                 |
| $E7    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $E8    | INX      | Implied          | 1     | 2                                         | Incrementa el registro X.                                             |
| $E9    | SBC      | Immediate        | 2     | 2                                         | Resta con acarreo con un valor inmediato.                            |
| $EA    | NOP      | Implied          | 1     | 2                                         | No realiza ninguna operación.                                         |
| $EB    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $EC    | CPX      | Absolute         | 3     | 4                                         | Compara el registro X con un valor en dirección absoluta.             |
| $ED    | SBC      | Absolute         | 3     | 4                                         | Resta con acarreo en dirección absoluta.                             |
| $EE    | INC      | Absolute         | 3     | 6                                         | Incrementa el valor en dirección absoluta.                           |
| $EF    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $F0    | BEQ      | Relative         | 2     | 2+ if branch taken, +2 if page crossed     | Salta si el flag cero es 1 (bifurcación).                             |
| $F1    | SBC      | (Indirect),Y    | 2     | 5+ if page crossed                         | Resta con acarreo usando indirección post-indexada en Y.              |
| $F2    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $F3    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $F4    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $F5    | SBC      | Zero Page,X      | 2     | 4                                         | Resta con acarreo en cero página con índice X.                        |
| $F6    | INC      | Zero Page,X      | 2     | 6                                         | Incrementa el valor en cero página con índice X.                     |
| $F7    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $F8    | SED      | Implied          | 1     | 2                                         | Establece el flag decimal.                                            |
| $F9    | SBC      | Absolute,Y       | 3     | 4+ if page crossed                         | Resta con acarreo en dirección absoluta con índice Y.                |
| $FA    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $FB    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $FC    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
| $FD    | SBC      | Absolute,X       | 3     | 4+ if page crossed                         | Resta con acarreo en dirección absoluta con índice X.                |
| $FE    | INC      | Absolute,X       | 3     | 7                                         | Incrementa el valor en dirección absoluta con índice X.              |
| $FF    | ILLEGAL  |                  |       |                                           | Instrucción no oficial.                                             |
